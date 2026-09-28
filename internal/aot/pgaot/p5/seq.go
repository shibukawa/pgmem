package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSeqScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	F_MemoryContextReset(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScan[0]))
		if v13 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				if v19 == int32(0) {
					v22 = F_ScanRelIsReadOnly(m, l0)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScan[1]))
						if v25 != 0 {
							v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecSeqScan[2])))
							if v27&int32(1) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_ExecSeqScan_0), int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ExecSeqScan_1), int32(931), int32(_a_F_ExecSeqScan_2))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								v34 = int32(0)
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
								if v22 != 0 {
									v44 = int32(1473)
								} else {
									v44 = int32(449)
								}
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+188))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
								v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, v33, v34, v34, v34, v37<<(uint(int32(7))%32)&int32(2048)|v44)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v48
									v52 = v48
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
									v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v16)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 != 0 {
											v62 = v16
										} else {
											v62 = int32(0)
										}
										return v62
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							v34 = int32(0)
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
							if v22 != 0 {
								v44 = int32(1473)
							} else {
								v44 = int32(449)
							}
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+188))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
							v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, v33, v34, v34, v34, v37<<(uint(int32(7))%32)&int32(2048)|v44)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v48
								v52 = v48
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
								v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v16)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									if v60 != 0 {
										v62 = v16
									} else {
										v62 = int32(0)
									}
									return v62
								}
							}
						}
					}
				} else {
					v52 = v19
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
					v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v16)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v60 != 0 {
							v62 = v16
						} else {
							v62 = int32(0)
						}
						return v62
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
			if v19 == int32(0) {
				v22 = F_ScanRelIsReadOnly(m, l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScan[1]))
					if v25 != 0 {
						v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecSeqScan[2])))
						if v27&int32(1) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_ExecSeqScan_0), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ExecSeqScan_1), int32(931), int32(_a_F_ExecSeqScan_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							v34 = int32(0)
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
							if v22 != 0 {
								v44 = int32(1473)
							} else {
								v44 = int32(449)
							}
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+188))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
							v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, v33, v34, v34, v34, v37<<(uint(int32(7))%32)&int32(2048)|v44)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v48
								v52 = v48
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
								v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v16)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									if v60 != 0 {
										v62 = v16
									} else {
										v62 = int32(0)
									}
									return v62
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						v34 = int32(0)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
						if v22 != 0 {
							v44 = int32(1473)
						} else {
							v44 = int32(449)
						}
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+188))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
						v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, v33, v34, v34, v34, v37<<(uint(int32(7))%32)&int32(2048)|v44)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v48
							v52 = v48
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
							v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v16)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								if v60 != 0 {
									v62 = v16
								} else {
									v62 = int32(0)
								}
								return v62
							}
						}
					}
				}
			} else {
				v52 = v19
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
				v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v16)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					if v60 != 0 {
						v62 = v16
					} else {
						v62 = int32(0)
					}
					return v62
				}
			}
		}
	}
}
func F_fill_seq_with_data(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_fill_seq_fork_with_data(m, l0, l1, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+118)))
		if v13 == int32(117) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v18
			v21 = F_smgropen(m, v7, int32(-1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_smgrcreate(m, v21, int32(3), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_log_smgrcreate(m, l0, int32(3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_fill_seq_fork_with_data(m, l0, l1, int32(3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_FlushRelationBuffers(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								F_smgrclose(m, v21)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_seq_desc(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	Fn14372(m, l0, l1, int32(_a_F_seq_desc_0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
