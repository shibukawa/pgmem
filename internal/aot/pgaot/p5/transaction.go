package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsTransactionState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsTransactionState[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	return base.B2i32(v3 == int32(2))
}
func F_TransactionIdAbortTree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	F_TransactionIdSetTreeStatus(m, l0, l1, l2, int32(2), int64(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[0]))
	if v10 == l0 {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[1]))
		v32 = v13
		v33 = int32(1)
		switch v32 - v33 {
		case 0:
			v75 = v33
			m.G0 = v7 + int32(16)
			return v75
		default:
			v75 = int32(0)
			m.G0 = v7 + int32(16)
			return v75
		case 2:
			v36 = int32(0)
			v37 = int32(3)
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[2]))
			if base.B2i32(base.Ui32(l0) < base.Ui32(v37))|base.B2i32(base.Ui32(v40) < base.Ui32(v37)) == v36 {
				if int32(0) <= l0-v40 {
					v50 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if v50 == int32(0) {
							v54 = int32(0)
							v57 = F_errstart(m, int32(19), v54)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 == int32(0) {
									v75 = v54
									m.G0 = v7 + int32(16)
									return v75
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										}
									}
								}
							}
						} else {
							v70 = F_TransactionIdDidCommit(m, v50)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v75 = v70
								m.G0 = v7 + int32(16)
								return v75
							}
						}
					}
				} else {
					v75 = v36
					m.G0 = v7 + int32(16)
					return v75
				}
			} else {
				if base.Ui32(l0) < base.Ui32(v40) {
					v75 = v36
					m.G0 = v7 + int32(16)
					return v75
				} else {
					v50 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if v50 == int32(0) {
							v54 = int32(0)
							v57 = F_errstart(m, int32(19), v54)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 == int32(0) {
									v75 = v54
									m.G0 = v7 + int32(16)
									return v75
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										}
									}
								}
							}
						} else {
							v70 = F_TransactionIdDidCommit(m, v50)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v75 = v70
								m.G0 = v7 + int32(16)
								return v75
							}
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			if l0 == int32(0) {
				v75 = int32(0)
			} else {
				v75 = int32(1)
			}
			m.G0 = v7 + int32(16)
			return v75
		} else {
			v21 = F_TransactionIdGetStatus(m, l0, v7+int32(8))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				switch v21 {
				case 0, 3:
					v32 = v21
				default:
					*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[1])) = v21
					*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[0])) = l0
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
					*(*int64)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[3])) = v30
					v32 = v21
				}
				v33 = int32(1)
				switch v32 - v33 {
				case 0:
					v75 = v33
					m.G0 = v7 + int32(16)
					return v75
				default:
					v75 = int32(0)
					m.G0 = v7 + int32(16)
					return v75
				case 2:
					v36 = int32(0)
					v37 = int32(3)
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[2]))
					if base.B2i32(base.Ui32(l0) < base.Ui32(v37))|base.B2i32(base.Ui32(v40) < base.Ui32(v37)) == v36 {
						if int32(0) <= l0-v40 {
							v50 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v54 = int32(0)
									v57 = F_errstart(m, int32(19), v54)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if v57 == int32(0) {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v75 = v54
													m.G0 = v7 + int32(16)
													return v75
												}
											}
										}
									}
								} else {
									v70 = F_TransactionIdDidCommit(m, v50)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = v70
										m.G0 = v7 + int32(16)
										return v75
									}
								}
							}
						} else {
							v75 = v36
							m.G0 = v7 + int32(16)
							return v75
						}
					} else {
						if base.Ui32(l0) < base.Ui32(v40) {
							v75 = v36
							m.G0 = v7 + int32(16)
							return v75
						} else {
							v50 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v54 = int32(0)
									v57 = F_errstart(m, int32(19), v54)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if v57 == int32(0) {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v75 = v54
													m.G0 = v7 + int32(16)
													return v75
												}
											}
										}
									}
								} else {
									v70 = F_TransactionIdDidCommit(m, v50)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = v70
										m.G0 = v7 + int32(16)
										return v75
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
func F_TransactionIdGetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = base.I32_div_u_s(l0, int32(819))
	if l0 != 0 {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
			if l2 != 0 {
				v91 = v4
				v94 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v94)
				v132 = v91
			} else {
				v132 = v4
			}
			m.G0 = v11 + int32(16)
			return v132
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[0]))
			v24 = F_LWLockAcquire(m, v20+int32(_a_F_TransactionIdGetCommitTsData_0), int32(1))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[1]))
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
				if v30 == int32(0) {
					F_error_commit_ts_disabled(m)
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if l0 == v33 {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = v35
						if l2 != 0 {
							v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v37)
						} else {
						}
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[0]))
						F_LWLockRelease(m, v40+int32(_a_F_TransactionIdGetCommitTsData_0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							v132 = base.B2i32(v45 != int64(0))
							m.G0 = v11 + int32(16)
							return v132
						}
					} else {
						v48 = int32(0)
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[2]))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[0]))
						F_LWLockRelease(m, v54+int32(_a_F_TransactionIdGetCommitTsData_0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v52 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
								if l2 == int32(0) {
									v132 = v48
								} else {
									v91 = v48
									v94 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v94)
									v132 = v91
								}
								m.G0 = v11 + int32(16)
								return v132
							} else {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v52))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
									v72 = base.B2i32(base.Ui32(l0) < base.Ui32(v52))
								} else {
									v72 = int32(base.Ui32(l0-v52) >> (uint(int32(31)) % 32))
								}
								if v72 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
									if l2 == int32(0) {
										v132 = v48
									} else {
										v91 = v48
										v94 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v94)
										v132 = v91
									}
									m.G0 = v11 + int32(16)
									return v132
								} else {
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v51)) == int32(0) {
										v84 = base.B2i32(base.Ui32(v51) < base.Ui32(l0))
									} else {
										v84 = int32(base.Ui32(v51-l0) >> (uint(int32(31)) % 32))
									}
									if v84 == int32(0) {
										v98 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_TransactionIdGetCommitTsData_1), base.I64_extend_i32_u(v14), l0)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[3]))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v98<<(uint(int32(2))%32))))
											v112 = v106 + (l0-v14*int32(819))*int32(10)
											v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+8)))
											v114 = *(*int64)(unsafe.Add(mBase, uint32(v112)))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v114
											if l2 != 0 {
												*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v113)
											} else {
											}
											v118 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[3]))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
											v121 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[4])))
											v122 = base.I32_rem_u_s(v14, v121)
											F_LWLockRelease(m, v119+v122<<(uint(int32(7))%32))
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												v128 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
												v132 = base.B2i32(v128 != int64(0))
												m.G0 = v11 + int32(16)
												return v132
											}
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
										if l2 == int32(0) {
											v132 = v48
										} else {
											v91 = v48
											v94 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v94)
											v132 = v91
										}
										m.G0 = v11 + int32(16)
										return v132
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v142 = m.ExcPending
		if v142 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
				F_errmsg(m, int32(_a_F_TransactionIdGetCommitTsData_2), v11)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_TransactionIdGetCommitTsData_3), int32(287), int32(_a_F_TransactionIdGetCommitTsData_4))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
