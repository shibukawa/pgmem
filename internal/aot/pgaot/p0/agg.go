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
		switch v7 - int32(435) {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+376))
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
func F_ExecAggCopyTransValue(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int64, l5 int32) int64 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	if l3 != 0 {
		v42 = int64(0)
		if l5 == int32(0) {
			v45 = base.I32_wrap_i64(l4)
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
			if v46 != int32(_a_F_ExecAggCopyTransValue_0) {
				F_pfree(m, v45)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					return v42
				}
			} else {
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
				if v49 != int32(1) {
					F_pfree(m, v45)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						return v42
					}
				} else {
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
					if v52 != int32(3) {
						F_pfree(m, v45)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
						} else {
							return v42
						}
					} else {
						F_DeleteExpandedObject(m, l4)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							return v42
						}
					}
				}
			}
		} else {
			return v42
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecAggCopyTransValue[0])) = v10
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
		if v12 != int32(_a_F_ExecAggCopyTransValue_0) {
			v33 = v12
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+191)))
			v36 = F_datumCopy(m, l2, v34, base.I32_extend16_s(v33))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				v42 = v36
				if l5 == int32(0) {
					v45 = base.I32_wrap_i64(l4)
					v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
					if v46 != int32(_a_F_ExecAggCopyTransValue_0) {
						F_pfree(m, v45)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
						} else {
							return v42
						}
					} else {
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
						if v49 != int32(1) {
							F_pfree(m, v45)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								return v42
							}
						} else {
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
							if v52 != int32(3) {
								F_pfree(m, v45)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									return v42
								}
							} else {
								F_DeleteExpandedObject(m, l4)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									return v42
								}
							}
						}
					}
				} else {
					return v42
				}
			}
		} else {
			v15 = base.I32_wrap_i64(l2)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v16 != int32(1) {
				v33 = int32(_a_F_ExecAggCopyTransValue_0)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+191)))
				v36 = F_datumCopy(m, l2, v34, base.I32_extend16_s(v33))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					v42 = v36
					if l5 == int32(0) {
						v45 = base.I32_wrap_i64(l4)
						v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
						if v46 != int32(_a_F_ExecAggCopyTransValue_0) {
							F_pfree(m, v45)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								return v42
							}
						} else {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							if v49 != int32(1) {
								F_pfree(m, v45)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									return v42
								}
							} else {
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
								if v52 != int32(3) {
									F_pfree(m, v45)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										return v42
									}
								} else {
									F_DeleteExpandedObject(m, l4)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
									} else {
										return v42
									}
								}
							}
						}
					} else {
						return v42
					}
				}
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
				if v21 != int32(3) {
					v33 = int32(_a_F_ExecAggCopyTransValue_0)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+191)))
					v36 = F_datumCopy(m, l2, v34, base.I32_extend16_s(v33))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v42 = v36
						if l5 == int32(0) {
							v45 = base.I32_wrap_i64(l4)
							v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
							if v46 != int32(_a_F_ExecAggCopyTransValue_0) {
								F_pfree(m, v45)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									return v42
								}
							} else {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								if v49 != int32(1) {
									F_pfree(m, v45)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										return v42
									}
								} else {
									v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
									if v52 != int32(3) {
										F_pfree(m, v45)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											return v42
										}
									} else {
										F_DeleteExpandedObject(m, l4)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int64(0)
										} else {
											return v42
										}
									}
								}
							}
						} else {
							return v42
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l2))+2))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAggCopyTransValue[0]))
					if v27 == v29 {
						v42 = l2
						if l5 == int32(0) {
							v45 = base.I32_wrap_i64(l4)
							v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
							if v46 != int32(_a_F_ExecAggCopyTransValue_0) {
								F_pfree(m, v45)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									return v42
								}
							} else {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								if v49 != int32(1) {
									F_pfree(m, v45)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										return v42
									}
								} else {
									v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
									if v52 != int32(3) {
										F_pfree(m, v45)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											return v42
										}
									} else {
										F_DeleteExpandedObject(m, l4)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int64(0)
										} else {
											return v42
										}
									}
								}
							}
						} else {
							return v42
						}
					} else {
						v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
						v33 = v31
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+191)))
						v36 = F_datumCopy(m, l2, v34, base.I32_extend16_s(v33))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v42 = v36
							if l5 == int32(0) {
								v45 = base.I32_wrap_i64(l4)
								v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+188)))
								if v46 != int32(_a_F_ExecAggCopyTransValue_0) {
									F_pfree(m, v45)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										return v42
									}
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
									if v49 != int32(1) {
										F_pfree(m, v45)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											return v42
										}
									} else {
										v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
										if v52 != int32(3) {
											F_pfree(m, v45)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int64(0)
											} else {
												return v42
											}
										} else {
											F_DeleteExpandedObject(m, l4)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int64(0)
											} else {
												return v42
											}
										}
									}
								}
							} else {
								return v42
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	v16 = F_palloc0(m, int32(112))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1584842932533)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		v25 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v24
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v28 == int32(1) {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			v33 = v31
		} else {
			v33 = int32(0)
		}
		v34 = int32(1)
		v35 = v33 & v34
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v35)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v37
		if l4 != v34 {
			v50 = int32(0)
			*(*float64)(unsafe.Add(mBase, uint32(v16)+88)) = l9
			*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v50
			if l8 != 0 {
				v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l8)+32)))
				v58 = v56
			} else {
				v58 = int64(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = l7
			*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = l6
			*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v58
			if l6 != 0 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
				v64 = v62
			} else {
				v64 = int32(0)
			}
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v68 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+32))
			F_cost_agg(m, v16, l0, l4, l8, v64, l9, l7, v65, v66, v67, v68, base.F64_convert_i32_s(v70))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				v75 = *(*float64)(unsafe.Add(mBase, uint32(v16)+48))
				*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_add(v74, v75)
				v78 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
				v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
				v80 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
				v82 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = base.F64_add(v78, base.F64_add(base.F64_mul(v79, v80), v82))
				return v16
			}
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
			if v41 != 0 {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
				v44 = v42
			} else {
				v44 = int32(0)
			}
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
			if v44 <= v45 {
				v50 = v41
				*(*float64)(unsafe.Add(mBase, uint32(v16)+88)) = l9
				*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v50
				if l8 != 0 {
					v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l8)+32)))
					v58 = v56
				} else {
					v58 = int64(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = l6
				*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v58
				if l6 != 0 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
					v64 = v62
				} else {
					v64 = int32(0)
				}
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
				v66 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
				v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
				v68 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+32))
				F_cost_agg(m, v16, l0, l4, l8, v64, l9, l7, v65, v66, v67, v68, base.F64_convert_i32_s(v70))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					v75 = *(*float64)(unsafe.Add(mBase, uint32(v16)+48))
					*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_add(v74, v75)
					v78 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
					v80 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
					v82 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = base.F64_add(v78, base.F64_add(base.F64_mul(v79, v80), v82))
					return v16
				}
			} else {
				v47 = F_list_copy_head(m, v41, v45)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v50 = v47
					*(*float64)(unsafe.Add(mBase, uint32(v16)+88)) = l9
					*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v50
					if l8 != 0 {
						v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l8)+32)))
						v58 = v56
					} else {
						v58 = int64(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = l6
					*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v58
					if l6 != 0 {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
						v64 = v62
					} else {
						v64 = int32(0)
					}
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
					v66 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
					v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
					v68 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+32))
					F_cost_agg(m, v16, l0, l4, l8, v64, l9, l7, v65, v66, v67, v68, base.F64_convert_i32_s(v70))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v16)+48))
						*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_add(v74, v75)
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v80 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
						v82 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = base.F64_add(v78, base.F64_add(base.F64_mul(v79, v80), v82))
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
			F_errmsg_internal(m, int32(_a_F_get_agg_combine_expr_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_get_agg_combine_expr_1), int32(_a_F_get_agg_combine_expr_2), int32(_a_F_get_agg_combine_expr_3))
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
