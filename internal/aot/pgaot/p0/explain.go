package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExplainDummyGroup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	switch v4 - int32(1) {
	case 0:
		F_ExplainXMLTag(m, l0, int32(2), l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	case 1:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_appendStringInfoChar(m, v13, int32(44))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_appendStringInfoChar(m, v19, int32(10))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					F_appendStringInfoSpaces(m, v23, v24<<(uint(int32(1))%32))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_escape_json(m, v29, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_appendStringInfoChar(m, v19, int32(10))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				F_appendStringInfoSpaces(m, v23, v24<<(uint(int32(1))%32))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_escape_json(m, v29, l0)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	case 2:
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
		if v34 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(1)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_appendStringInfoString(m, v49, int32(747365))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_escape_json(m, v53, l0)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_appendStringInfoChar(m, v39, int32(10))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				F_appendStringInfoSpaces(m, v43, v44<<(uint(int32(1))%32))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_appendStringInfoString(m, v49, int32(747365))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_escape_json(m, v53, l0)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	default:
		return
	}
}
func F_ExplainProperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	switch v12 {
	case 0:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v14-int32(1)))))
			if v19 != int32(10) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if l1 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
					F_appendStringInfo(m, v27, int32(749447), v10+int32(16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
					F_appendStringInfo(m, v27, int32(749641), v10)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
				F_appendStringInfoSpaces(m, v13, v22<<(uint(int32(1))%32))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
						F_appendStringInfo(m, v27, int32(749447), v10+int32(16))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
						F_appendStringInfo(m, v27, int32(749641), v10)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
			F_appendStringInfoSpaces(m, v13, v22<<(uint(int32(1))%32))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if l1 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
					F_appendStringInfo(m, v27, int32(749447), v10+int32(16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
					F_appendStringInfo(m, v27, int32(749641), v10)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				}
			}
		}
	case 1:
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
		F_appendStringInfoSpaces(m, v41, v42<<(uint(int32(1))%32))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			F_ExplainXMLTag(m, l0, int32(4), l4)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				v50 = F_escape_xml(m, l2)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					F_appendStringInfoString(m, v52, v50)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_pfree(m, v50)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_ExplainXMLTag(m, l0, int32(5), l4)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								F_appendStringInfoChar(m, v60, int32(10))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v10 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	case 2:
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
		if v66 != 0 {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			F_appendStringInfoChar(m, v67, int32(44))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				F_appendStringInfoChar(m, v73, int32(10))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
					F_appendStringInfoSpaces(m, v77, v78<<(uint(int32(1))%32))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						F_escape_json(m, v83, l0)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							F_appendStringInfoString(m, v86, int32(747356))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if l3 != 0 {
									F_appendStringInfoString(m, v90, l2)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										m.G0 = v10 + int32(48)
										return
									}
								} else {
									F_escape_json(m, v90, l2)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										m.G0 = v10 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(1)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			F_appendStringInfoChar(m, v73, int32(10))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
				F_appendStringInfoSpaces(m, v77, v78<<(uint(int32(1))%32))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					F_escape_json(m, v83, l0)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						F_appendStringInfoString(m, v86, int32(747356))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if l3 != 0 {
								F_appendStringInfoString(m, v90, l2)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									m.G0 = v10 + int32(48)
									return
								}
							} else {
								F_escape_json(m, v90, l2)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									m.G0 = v10 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v95 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
		if v97 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(1)
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
			F_appendStringInfo(m, v112, int32(746999), v10+int32(32))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if l3 != 0 {
					F_appendStringInfoString(m, v119, l2)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				} else {
					F_escape_json(m, v119, l2)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				}
			}
		} else {
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			F_appendStringInfoChar(m, v102, int32(10))
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return
			} else {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
				F_appendStringInfoSpaces(m, v106, v107<<(uint(int32(1))%32))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
					F_appendStringInfo(m, v112, int32(746999), v10+int32(32))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if l3 != 0 {
							F_appendStringInfoString(m, v119, l2)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						} else {
							F_escape_json(m, v119, l2)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				}
			}
		}
	default:
		m.G0 = v10 + int32(48)
		return
	}
}
func F_ExplainPropertyBool(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	if l1 != 0 {
		v7 = int32(345210)
	} else {
		v7 = int32(362226)
	}
	F_ExplainProperty(m, l0, int32(0), v7, int32(1), l2)
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_ExplainPropertyFloat(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
	v14 = F_psprintf(m, int32(340833), v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		F_ExplainProperty(m, l0, l1, v14, int32(1), l4)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_pfree(m, v14)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_ExplainPropertyList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	switch v11 {
	case 0:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L3
	case 3:
		goto L2
	default:
		goto L1
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return
L2:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v207 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L3:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v135 != 0 {
		goto L40
	} else {
		goto L41
	}
L4:
	;
	F_ExplainXMLTag(m, l0, int32(0), l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L25
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v13 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_appendStringInfo(m, v27, int32(746999), v9)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L13
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v13-int32(1)))))
	if v18 != int32(10) {
		v27 = v12
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_appendStringInfoSpaces(m, v12, v21<<(uint(int32(1))%32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	return
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v27 = v26
	goto L6
L13:
	;
	if l1 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v74, int32(10))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L24
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	F_appendStringInfoString(m, v38, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 <= int32(1) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(1)
	goto L19
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v53, int32(747599))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52+v49<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v57, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v65 = v49 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v65 < v66 {
		v49 = v65
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	goto L1
L25:
	;
	if l1 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_ExplainXMLTag(m, l0, int32(1), l2)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L11
	} else {
		goto L38
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v83 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v89 = int32(0)
	goto L29
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_appendStringInfoSpaces(m, v93, v94<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v101, int32(547338))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v92+v89<<(uint(int32(2))%32))))
	v109 = F_escape_xml(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v111, v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	F_pfree(m, v109)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v116, int32(755102))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v121 = v89 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v121 < v122 {
		v89 = v121
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	goto L1
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v142, int32(10))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L11
	} else {
		goto L44
	}
L40:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v136, int32(44))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(1)
	goto L39
L43:
	;
	goto L39
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_appendStringInfoSpaces(m, v146, v147<<(uint(int32(1))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_escape_json(m, v152, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v155, int32(509148))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	if l1 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v201, int32(93))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L58
	}
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v161 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_escape_json(m, v164, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v169 = int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v170 <= v169 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v176 = v169
	goto L53
L53:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v180, int32(747599))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L55
	}
L54:
	;
	goto L48
L55:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179+v176<<(uint(int32(2))%32))))
	F_escape_json(m, v184, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v192 = v176 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v192 < v193 {
		v176 = v192
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L1
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	F_appendStringInfo(m, v222, int32(746999), v9+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L11
	} else {
		goto L65
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = int32(1)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v212, int32(10))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_appendStringInfoSpaces(m, v216, v217<<(uint(int32(1))%32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v231 <= int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v238 = int32(0)
	goto L68
L68:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v242, int32(10))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L11
	} else {
		goto L70
	}
L69:
	;
	goto L1
L70:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_appendStringInfoSpaces(m, v246, v247<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v254, int32(747365))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v241+v238<<(uint(int32(2))%32))))
	F_escape_json(m, v258, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v266 = v238 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v266 < v267 {
		v238 = v266
		goto L68
	} else {
		goto L74
	}
L74:
	;
	goto L69
}
func F_ExplainSeparatePlans(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_appendStringInfoChar(m, v5, int32(10))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
