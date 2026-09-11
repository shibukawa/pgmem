package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EvalPlanQual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	F_EvalPlanQualBegin(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = v11 + l2<<(uint(int32(2))%32) - int32(4)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 == int32(0) {
			v20 = int32(_a_F_EvalPlanQual_0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
			*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v24
			v28 = F_table_slot_create(m, l1, l0+int32(12))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v28
				*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v21
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v34 = v33
				if l3 != v34 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
					m.T0[v37].(func(*base.Module, int32, int32))(m, v34, l3)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v41 = l2 - int32(1)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v44 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v41+v42))) = uint8(v44)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						*(*uint8)(unsafe.Add(mBase, uint32(v46+v41))) = uint8(v44)
						v50 = int32(_a_F_EvalPlanQual_0)
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0]))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
						*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v54
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
						if v57 != 0 {
							F_ExecReScan(m, v56)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
								v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
									if v61 == int32(0) {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									} else {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
										if v67&int32(2) != 0 {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
											m.T0[v75].(func(*base.Module, int32))(m, v34)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												v80 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
												return v61
											}
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
											m.T0[v71].(func(*base.Module, int32))(m, v61)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
												m.T0[v75].(func(*base.Module, int32))(m, v34)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													v80 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
													return v61
												}
											}
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
								if v61 == int32(0) {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
									if v67&int32(2) != 0 {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
										m.T0[v71].(func(*base.Module, int32))(m, v61)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
											m.T0[v75].(func(*base.Module, int32))(m, v34)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												v80 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
												return v61
											}
										}
									}
								}
							}
						}
					}
				} else {
					v41 = l2 - int32(1)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v44 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v41+v42))) = uint8(v44)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*uint8)(unsafe.Add(mBase, uint32(v46+v41))) = uint8(v44)
					v50 = int32(_a_F_EvalPlanQual_0)
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0]))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
					*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
					if v57 != 0 {
						F_ExecReScan(m, v56)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
								if v61 == int32(0) {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
									if v67&int32(2) != 0 {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
										m.T0[v71].(func(*base.Module, int32))(m, v61)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
											m.T0[v75].(func(*base.Module, int32))(m, v34)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												v80 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
												return v61
											}
										}
									}
								}
							}
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
						v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
							if v61 == int32(0) {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								m.T0[v75].(func(*base.Module, int32))(m, v34)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									v80 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
									return v61
								}
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
								if v67&int32(2) != 0 {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
									m.T0[v71].(func(*base.Module, int32))(m, v61)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v34 = v17
			if l3 != v34 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
				m.T0[v37].(func(*base.Module, int32, int32))(m, v34, l3)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = l2 - int32(1)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v44 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v41+v42))) = uint8(v44)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*uint8)(unsafe.Add(mBase, uint32(v46+v41))) = uint8(v44)
					v50 = int32(_a_F_EvalPlanQual_0)
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0]))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
					*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
					if v57 != 0 {
						F_ExecReScan(m, v56)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
								if v61 == int32(0) {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
									if v67&int32(2) != 0 {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
										m.T0[v71].(func(*base.Module, int32))(m, v61)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
											m.T0[v75].(func(*base.Module, int32))(m, v34)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												v80 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
												return v61
											}
										}
									}
								}
							}
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
						v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
							if v61 == int32(0) {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								m.T0[v75].(func(*base.Module, int32))(m, v34)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									v80 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
									return v61
								}
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
								if v67&int32(2) != 0 {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
									m.T0[v71].(func(*base.Module, int32))(m, v61)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									}
								}
							}
						}
					}
				}
			} else {
				v41 = l2 - int32(1)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v44 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v41+v42))) = uint8(v44)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*uint8)(unsafe.Add(mBase, uint32(v46+v41))) = uint8(v44)
				v50 = int32(_a_F_EvalPlanQual_0)
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0]))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
				*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
				if v57 != 0 {
					F_ExecReScan(m, v56)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
						v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
							if v61 == int32(0) {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								m.T0[v75].(func(*base.Module, int32))(m, v34)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									v80 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
									return v61
								}
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
								if v67&int32(2) != 0 {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
									m.T0[v71].(func(*base.Module, int32))(m, v61)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
										m.T0[v75].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
											return v61
										}
									}
								}
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
					v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v56)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQual[0])) = v51
						if v61 == int32(0) {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
							m.T0[v75].(func(*base.Module, int32))(m, v34)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								v80 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
								return v61
							}
						} else {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
							if v67&int32(2) != 0 {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								m.T0[v75].(func(*base.Module, int32))(m, v34)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									v80 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
									return v61
								}
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
								m.T0[v71].(func(*base.Module, int32))(m, v61)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									m.T0[v75].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										v80 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))) = uint8(v80)
										return v61
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
