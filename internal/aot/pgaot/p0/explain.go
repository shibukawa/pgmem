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
			F_appendStringInfoString(m, v49, int32(_a_F_ExplainDummyGroup_0))
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
					F_appendStringInfoString(m, v49, int32(_a_F_ExplainDummyGroup_0))
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
					F_appendStringInfo(m, v27, int32(_a_F_ExplainProperty_0), v10+int32(16))
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
					F_appendStringInfo(m, v27, int32(_a_F_ExplainProperty_1), v10)
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
						F_appendStringInfo(m, v27, int32(_a_F_ExplainProperty_0), v10+int32(16))
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
						F_appendStringInfo(m, v27, int32(_a_F_ExplainProperty_1), v10)
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
					F_appendStringInfo(m, v27, int32(_a_F_ExplainProperty_0), v10+int32(16))
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
					F_appendStringInfo(m, v27, int32(_a_F_ExplainProperty_1), v10)
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
							F_appendStringInfoString(m, v86, int32(_a_F_ExplainProperty_2))
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
						F_appendStringInfoString(m, v86, int32(_a_F_ExplainProperty_2))
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
			F_appendStringInfo(m, v112, int32(_a_F_ExplainProperty_3), v10+int32(32))
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
					F_appendStringInfo(m, v112, int32(_a_F_ExplainProperty_3), v10+int32(32))
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
		v7 = int32(_a_F_ExplainPropertyBool_0)
	} else {
		v7 = int32(_a_F_ExplainPropertyBool_1)
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
	v14 = F_psprintf(m, int32(_a_F_ExplainPropertyFloat_0), v9)
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
	F_appendStringInfo(m, v27, int32(_a_F_ExplainPropertyList_0), v9)
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
	F_appendStringInfoString(m, v53, int32(_a_F_ExplainPropertyList_1))
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
	F_appendStringInfoString(m, v101, int32(_a_F_ExplainPropertyList_2))
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
	F_appendStringInfoString(m, v116, int32(_a_F_ExplainPropertyList_3))
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
	F_appendStringInfoString(m, v155, int32(_a_F_ExplainPropertyList_4))
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
	F_appendStringInfoString(m, v180, int32(_a_F_ExplainPropertyList_1))
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
	F_appendStringInfo(m, v222, int32(_a_F_ExplainPropertyList_0), v9+int32(16))
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
	F_appendStringInfoString(m, v254, int32(_a_F_ExplainPropertyList_5))
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
func F_explain_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[0]))
	if v280 != 0 {
		goto L60
	} else {
		goto L61
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[1]))
	if v17 < int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[2]))
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[3])))
	if v23 != int32(1) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[4])))
	if v27 != int32(1) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	v30 = int32(_a_F_explain_ExecutorEnd_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[5]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[5])) = v34
	F_InstrEndLoop(m, v13)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v39 = *(*float64)(unsafe.Add(mBase, uint32(v38)+208))
	v41 = base.F64_mul(v39, float64(1000))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[1]))
	if base.F64_ge(v41, base.F64_convert_i32_s(v43)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[5])) = v31
	goto L1
L12:
	;
	v48 = F_NewExplainState(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[6])))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v55 = v51 & base.B2i32(v52 != int32(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+5)) = uint8(v55)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[7])))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)) = uint8(v58)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[8])))
	v62 = v55 & v61
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+7)) = uint8(v62)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[9])))
	v66 = v55 & v65
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+8)) = uint8(v66)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[10])))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+10)) = uint8(v55)
	v71 = v55 & v69
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+9)) = uint8(v71)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v74
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[12])))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)) = uint8(v77)
	F_ExplainBeginOutput(m, v48)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_ExplainPropertyText(m, int32(_a_F_explain_ExecutorEnd_1), v81, v48)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[13]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v87 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	F_ExplainPrintPlan(m, v48, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L27
	}
L20:
	;
	if v86 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)+28))
	if v92 <= int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v96 = F_BuildParamLogString(m, v87, int32(0), v86)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	if v96 == int32(0) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_ExplainPropertyText(m, int32(_a_F_explain_ExecutorEnd_5), v96, v48)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+5)))
	if v109 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+6)))
	if v118 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[15])))
	if v113 != int32(1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_ExplainPrintTriggers(m, v48, l0)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v121 = m.G0
	v123 = v121 - int32(48)
	m.G0 = v123
	v125 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v123)+40)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v123)+32)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v123)+24)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v123)+16)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = v125
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+176)))
	if v138&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	F_ExplainEndOutput(m, v48)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L47
	}
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+180))
	if v141 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	m.G0 = v123 + int32(48)
	goto L34
L38:
	;
	v143 = v141 + int32(8)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v144 + v145
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v123)+8))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v148 + v149
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v143)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+16)) = v152 + v153
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v143)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+24)) = v156 + v157
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v123)+32))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v143)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+32)) = v160 + v161
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v123)+40))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v143)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+40)) = v164 + v165
	goto L41
L39:
	;
	v169 = v137
	goto L40
L40:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+184))
	if v170 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v169 = v168
	goto L40
L42:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v171 + v172
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v123)+8))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v175 + v176
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v170)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+16)) = v179 + v180
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v170)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+24)) = v183 + v184
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v123)+32))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v170)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+32)) = v187 + v188
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v123)+40))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v170)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+40)) = v191 + v192
	goto L45
L43:
	;
	v196 = v169
	goto L44
L44:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+176))
	F_ExplainPrintJIT(m, v48, v197, v123)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v196 = v195
	goto L44
L46:
	;
	goto L37
L47:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v211 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[11]))
	if v230 == int32(2) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214+v211-int32(1)))))
	if v218 != int32(10) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v222 = v211 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v222
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v222+v214))) = uint8(v225)
	goto L48
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = int32(123)
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v235)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v243 = int32(125)
	*(*uint8)(unsafe.Add(mBase, uint32(v238+v239-int32(1)))) = uint8(v243)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorEnd[14]))
	v249 = F_errstart(m, v247, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	if v249 == int32(0) {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v254
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v41
	F_errmsg(m, int32(_a_F_explain_ExecutorEnd_2), v11)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	F_errhidestmt(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_explain_ExecutorEnd_3), int32(437), int32(_a_F_explain_ExecutorEnd_4))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	goto L11
L59:
	;
	m.G0 = v11 + int32(16)
	return
L60:
	;
	m.T0[v280].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_standard_ExecutorEnd(m, l0)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L64
	}
L63:
	;
	goto L59
L64:
	;
	goto L59
}
func F_explain_ExecutorFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = v2
	v17 = v12
	v18 = int32(-1)
	v20 = v2
	v21 = v2
	v22 = v2
	goto L1
L1:
	;
	if v18 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[0])) = v47
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[1])) = v48
	v98 = int32(_a_F_explain_ExecutorFinish_0)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[2])) = v99 - int32(1)
	m.G0 = v12 + int32(16)
	return
L3:
	;
	v27 = v17 - int32(160)
	m.G0 = v27
	v29 = int32(_a_F_explain_ExecutorFinish_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[2]))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[2])) = v30 + v31
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[0]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v12 + int32(12)
	goto L6
L4:
	;
	v44 = v16
	v45 = v17
	v46 = v20
	v47 = v21
	v48 = v22
	goto L5
L5:
	;
	goto L8
L6:
	;
	v44 = int32(0)
	v45 = v27
	v46 = v27
	v47 = v35
	v48 = v37
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v44 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v73 = int32(m.ExcTag)
	v74 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v73 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorFinish(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[0])) = v46
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[3]))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[1])) = v48
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[0])) = v47
	v63 = int32(_a_F_explain_ExecutorFinish_0)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorFinish[2])) = v64 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v54].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v78 = int32(v74)
	m.G0 = v45
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v12+int32(12) == v85 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v87
	goto L24
L23:
	;
	v88 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v81, v80)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v16 = v80
	v17 = v45
	v18 = v88
	v20 = v46
	v21 = v47
	v22 = v48
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_explain_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v20 = v4
	v21 = v14
	v22 = int32(-1)
	v24 = v4
	v25 = v4
	v26 = v4
	goto L1
L1:
	;
	if v22 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[0])) = v51
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[1])) = v52
	v102 = int32(_a_F_explain_ExecutorRun_0)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[2])) = v103 - int32(1)
	m.G0 = v14 + int32(16)
	return
L3:
	;
	v31 = v21 - int32(160)
	m.G0 = v31
	v33 = int32(_a_F_explain_ExecutorRun_0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[2]))
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[2])) = v34 + v35
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[0]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v14 + int32(12)
	goto L6
L4:
	;
	v48 = v20
	v49 = v21
	v50 = v24
	v51 = v25
	v52 = v26
	goto L5
L5:
	;
	goto L8
L6:
	;
	v48 = int32(0)
	v49 = v31
	v50 = v31
	v51 = v39
	v52 = v41
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v77 = int32(m.ExcTag)
	v78 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v77 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorRun(m, l0, l1, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[0])) = v50
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[3]))
	if v58 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[1])) = v52
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[0])) = v51
	v67 = int32(_a_F_explain_ExecutorRun_0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorRun[2])) = v68 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v58].(func(*base.Module, int32, int32, int64))(m, l0, l1, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v82 = int32(v78)
	m.G0 = v49
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v14+int32(12) == v89 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v92 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v92 = v91
	goto L24
L23:
	;
	v92 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v85, v84)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v20 = v84
	v21 = v49
	v22 = v92
	v24 = v50
	v25 = v51
	v26 = v52
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
