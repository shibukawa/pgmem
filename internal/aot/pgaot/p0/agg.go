package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AggCheckCallContext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 == v3 {
		v23 = int32(0)
		if l1 == v23 {
			v31 = v23
		} else {
			v26 = v23
			v27 = v3
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
			v31 = v27
		}
		return v31
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		switch v7 - int32(429) {
		case 0:
			if l1 == int32(0) {
				return int32(1)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+168))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
				v26 = v15
				v27 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
				v31 = v27
				return v31
			}
		case 1:
			if l1 == int32(0) {
				return int32(2)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+368))
				v26 = v21
				v27 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
				v31 = v27
				return v31
			}
		default:
			v23 = int32(0)
			if l1 == v23 {
				v31 = v23
			} else {
				v26 = v23
				v27 = v3
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
				v31 = v27
			}
			return v31
		}
	}
}
func F_ExecAggCopyTransValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	if l3 != 0 {
		v38 = int32(0)
		if l5 == int32(0) {
			v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
			if v41 != int32(65535) {
				F_pfree(m, l4)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					return v38
				}
			} else {
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
				if v44 != int32(1) {
					F_pfree(m, l4)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						return v38
					}
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
					if v47 != int32(3) {
						F_pfree(m, l4)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							return v38
						}
					} else {
						F_DeleteExpandedObject(m, l4)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							return v38
						}
					}
				}
			}
		} else {
			return v38
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
		if v12 != int32(65535) {
			v30 = v12
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
			v33 = F_datumCopy(m, l2, v31, base.I32_extend16_s(v30))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v38 = v33
				if l5 == int32(0) {
					v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
					if v41 != int32(65535) {
						F_pfree(m, l4)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							return v38
						}
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
						if v44 != int32(1) {
							F_pfree(m, l4)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								return v38
							}
						} else {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
							if v47 != int32(3) {
								F_pfree(m, l4)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									return v38
								}
							} else {
								F_DeleteExpandedObject(m, l4)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									return v38
								}
							}
						}
					}
				} else {
					return v38
				}
			}
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
			if v15 != int32(1) {
				v30 = int32(65535)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
				v33 = F_datumCopy(m, l2, v31, base.I32_extend16_s(v30))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v38 = v33
					if l5 == int32(0) {
						v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
						if v41 != int32(65535) {
							F_pfree(m, l4)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								return v38
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
							if v44 != int32(1) {
								F_pfree(m, l4)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									return v38
								}
							} else {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
								if v47 != int32(3) {
									F_pfree(m, l4)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										return v38
									}
								} else {
									F_DeleteExpandedObject(m, l4)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										return v38
									}
								}
							}
						}
					} else {
						return v38
					}
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
				if v20 != int32(3) {
					v30 = int32(65535)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
					v33 = F_datumCopy(m, l2, v31, base.I32_extend16_s(v30))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v38 = v33
						if l5 == int32(0) {
							v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
							if v41 != int32(65535) {
								F_pfree(m, l4)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									return v38
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
								if v44 != int32(1) {
									F_pfree(m, l4)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										return v38
									}
								} else {
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
									if v47 != int32(3) {
										F_pfree(m, l4)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											return v38
										}
									} else {
										F_DeleteExpandedObject(m, l4)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											return v38
										}
									}
								}
							}
						} else {
							return v38
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+2))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
					v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					if v25 == v27 {
						v38 = l2
						if l5 == int32(0) {
							v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
							if v41 != int32(65535) {
								F_pfree(m, l4)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									return v38
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
								if v44 != int32(1) {
									F_pfree(m, l4)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										return v38
									}
								} else {
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
									if v47 != int32(3) {
										F_pfree(m, l4)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											return v38
										}
									} else {
										F_DeleteExpandedObject(m, l4)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											return v38
										}
									}
								}
							}
						} else {
							return v38
						}
					} else {
						v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
						v30 = v29
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
						v33 = F_datumCopy(m, l2, v31, base.I32_extend16_s(v30))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v38 = v33
							if l5 == int32(0) {
								v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+184)))
								if v41 != int32(65535) {
									F_pfree(m, l4)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										return v38
									}
								} else {
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
									if v44 != int32(1) {
										F_pfree(m, l4)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											return v38
										}
									} else {
										v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
										if v47 != int32(3) {
											F_pfree(m, l4)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												return v38
											}
										} else {
											F_DeleteExpandedObject(m, l4)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												return v38
											}
										}
									}
								}
							} else {
								return v38
							}
						}
					}
				}
			}
		}
	}
}
func F_create_agg_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	v11 = int32(0)
	v16 = F_palloc0(m, int32(112))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1567663063347)
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v28 == int32(1) {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			v32 = v31
		} else {
			v32 = v11
		}
		v33 = int32(1)
		v34 = v32 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v34)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v36
		if l4 != v33 {
			v48 = v11
			*(*float64)(unsafe.Add(mBase, uint32(v16)+88)) = l9
			*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v48
			if l8 != 0 {
				v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l8)+32)))
				v57 = v55
			} else {
				v57 = int64(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = l7
			*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = l6
			*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v57
			if l6 != 0 {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
				v63 = v61
			} else {
				v63 = int32(0)
			}
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
			v65 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
			F_cost_agg(m, v16, l0, l4, l8, v63, l9, l7, v64, v65, v66, v67, base.F64_convert_i32_s(v69))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				v74 = *(*float64)(unsafe.Add(mBase, uint32(v16)+48))
				*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_add(v73, v74)
				v77 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
				v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
				v81 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = base.F64_add(v77, base.F64_add(base.F64_mul(v78, v79), v81))
				return v16
			}
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
			if v40 != 0 {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v43 = v41
			} else {
				v43 = int32(0)
			}
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
			if v43 <= v44 {
				v48 = v40
				*(*float64)(unsafe.Add(mBase, uint32(v16)+88)) = l9
				*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v48
				if l8 != 0 {
					v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l8)+32)))
					v57 = v55
				} else {
					v57 = int64(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = l6
				*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v57
				if l6 != 0 {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
					v63 = v61
				} else {
					v63 = int32(0)
				}
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
				v65 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
				v66 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
				v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
				F_cost_agg(m, v16, l0, l4, l8, v63, l9, l7, v64, v65, v66, v67, base.F64_convert_i32_s(v69))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					v74 = *(*float64)(unsafe.Add(mBase, uint32(v16)+48))
					*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_add(v73, v74)
					v77 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
					v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
					v81 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = base.F64_add(v77, base.F64_add(base.F64_mul(v78, v79), v81))
					return v16
				}
			} else {
				v46 = F_list_copy_head(m, v40, v44)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = v46
					*(*float64)(unsafe.Add(mBase, uint32(v16)+88)) = l9
					*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v48
					if l8 != 0 {
						v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l8)+32)))
						v57 = v55
					} else {
						v57 = int64(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = l6
					*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v57
					if l6 != 0 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
						v63 = v61
					} else {
						v63 = int32(0)
					}
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
					v65 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
					v66 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
					v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
					F_cost_agg(m, v16, l0, l4, l8, v63, l9, l7, v64, v65, v66, v67, base.F64_convert_i32_s(v69))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v74 = *(*float64)(unsafe.Add(mBase, uint32(v16)+48))
						*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_add(v73, v74)
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
						v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
						v81 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = base.F64_add(v77, base.F64_add(base.F64_mul(v78, v79), v81))
						return v16
					}
				}
			}
		}
	}
}
func F_get_agg_combine_expr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != int32(9) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(339430), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(494218), int32(11015), int32(207078))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v20 = int32(0)
		F_get_agg_expr_helper(m, l0, l1, l2, v20, v20, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			return
		}
	}
}
