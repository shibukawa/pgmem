package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ActivateCommitTs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
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
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[0]))
	if v7 == int32(0) {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
		v15 = F_LWLockAcquire(m, v11+int32(_a_F_ActivateCommitTs_0), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[2]))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
			F_LWLockRelease(m, v21+int32(_a_F_ActivateCommitTs_0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if v19 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[3]))
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[4]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					v32 = base.I32_div_u_s(v30, int32(819))
					v33 = base.I64_extend_i32_u(v32)
					*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v33
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
					v40 = F_LWLockAcquire(m, v36+int32(_a_F_ActivateCommitTs_0), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[4]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
						if v44 == int32(0) {
							v47 = F_ReadNextFullTransactionId(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[4]))
								v51 = base.I32_wrap_i64(v47)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+40)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(v50)+44)) = v51
								v58 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
								F_LWLockRelease(m, v58+int32(_a_F_ActivateCommitTs_0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = F_SimpleLruDoesPhysicalPageExist(m, int32(_a_F_ActivateCommitTs_1), v33)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										if v64 == int32(0) {
											v69 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[3]))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
											v72 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_ActivateCommitTs[5])))
											v73 = base.I32_rem_u_s(v32, v72)
											v76 = v70 + v73<<(uint(int32(7))%32)
											v78 = F_LWLockAcquire(m, v76, int32(0))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												v80 = int32(_a_F_ActivateCommitTs_1)
												v82 = F_SimpleLruZeroPage(m, v80, v33)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													F_SimpleLruWritePage(m, v80, v82)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														F_LWLockRelease(m, v76)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
															v94 = F_LWLockAcquire(m, v90+int32(_a_F_ActivateCommitTs_0), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return
															} else {
																v97 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[2]))
																v98 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
																v101 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
																F_LWLockRelease(m, v101+int32(_a_F_ActivateCommitTs_0))
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
											v94 = F_LWLockAcquire(m, v90+int32(_a_F_ActivateCommitTs_0), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[2]))
												v98 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
												v101 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
												F_LWLockRelease(m, v101+int32(_a_F_ActivateCommitTs_0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
							F_LWLockRelease(m, v58+int32(_a_F_ActivateCommitTs_0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = F_SimpleLruDoesPhysicalPageExist(m, int32(_a_F_ActivateCommitTs_1), v33)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									if v64 == int32(0) {
										v69 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[3]))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
										v72 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_ActivateCommitTs[5])))
										v73 = base.I32_rem_u_s(v32, v72)
										v76 = v70 + v73<<(uint(int32(7))%32)
										v78 = F_LWLockAcquire(m, v76, int32(0))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v80 = int32(_a_F_ActivateCommitTs_1)
											v82 = F_SimpleLruZeroPage(m, v80, v33)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												F_SimpleLruWritePage(m, v80, v82)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_LWLockRelease(m, v76)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
														v94 = F_LWLockAcquire(m, v90+int32(_a_F_ActivateCommitTs_0), int32(0))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															v97 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[2]))
															v98 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
															v101 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
															F_LWLockRelease(m, v101+int32(_a_F_ActivateCommitTs_0))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																return
															}
														}
													}
												}
											}
										}
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
										v94 = F_LWLockAcquire(m, v90+int32(_a_F_ActivateCommitTs_0), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[2]))
											v98 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
											v101 = *(*int32)(unsafe.Add(mBase, _c_F_ActivateCommitTs[1]))
											F_LWLockRelease(m, v101+int32(_a_F_ActivateCommitTs_0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												return
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
	}
}
func F_TS_execute_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_TS_execute_recurse[0]))
		if v17 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v20 == int32(1) {
					v24 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v90 = v24
						m.G0 = v10 + int32(16)
						return v90
					}
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					switch v26 - int32(1) {
					case 0:
						v29 = int32(1)
						if l2&v29 != 0 {
							v90 = v29
							m.G0 = v10 + int32(16)
							return v90
						} else {
							v34 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v34) < base.Ui32(int32(3)) {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F_TS_execute_recurse[1])))
									v90 = v84
								} else {
									v90 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v90
							}
						}
					case 1:
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v42 = F_TS_execute_recurse(m, l0+v38*int32(12), l1, l2, l3)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v42 == int32(0) {
								v90 = int32(0)
								m.G0 = v10 + int32(16)
								return v90
							} else {
								v48 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									switch v48 {
									case 0, 2:
										v90 = v48
									case 1:
										v90 = v42
									default:
										v90 = int32(0)
									}
									m.G0 = v10 + int32(16)
									return v90
								}
							}
						}
					case 2:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v55 = F_TS_execute_recurse(m, l0+v51*int32(12), l1, l2, l3)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 == int32(1) {
								v90 = int32(1)
								m.G0 = v10 + int32(16)
								return v90
							} else {
								v61 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									switch v61 {
									case 0:
										v90 = v55
									case 1, 2:
										v90 = v61
									default:
										v90 = int32(0)
									}
									m.G0 = v10 + int32(16)
									return v90
								}
							}
						}
					case 3:
						v64 = F_TS_phrase_execute(m, l0, l1, l2, l3, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							switch v64 {
							case 0, 1:
								v90 = v64
							case 2:
								v90 = l2 & int32(2)
							default:
								v90 = int32(0)
							}
							m.G0 = v10 + int32(16)
							return v90
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
							F_errmsg_internal(m, int32(_a_F_TS_execute_recurse_0), v10)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_TS_execute_recurse_1), int32(1969), int32(_a_F_TS_execute_recurse_2))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
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
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v20 == int32(1) {
				v24 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v90 = v24
					m.G0 = v10 + int32(16)
					return v90
				}
			} else {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				switch v26 - int32(1) {
				case 0:
					v29 = int32(1)
					if l2&v29 != 0 {
						v90 = v29
						m.G0 = v10 + int32(16)
						return v90
					} else {
						v34 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if base.Ui32(v34) < base.Ui32(int32(3)) {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F_TS_execute_recurse[1])))
								v90 = v84
							} else {
								v90 = int32(0)
							}
							m.G0 = v10 + int32(16)
							return v90
						}
					}
				case 1:
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v42 = F_TS_execute_recurse(m, l0+v38*int32(12), l1, l2, l3)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 == int32(0) {
							v90 = int32(0)
							m.G0 = v10 + int32(16)
							return v90
						} else {
							v48 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								switch v48 {
								case 0, 2:
									v90 = v48
								case 1:
									v90 = v42
								default:
									v90 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v90
							}
						}
					}
				case 2:
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v55 = F_TS_execute_recurse(m, l0+v51*int32(12), l1, l2, l3)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 == int32(1) {
							v90 = int32(1)
							m.G0 = v10 + int32(16)
							return v90
						} else {
							v61 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								switch v61 {
								case 0:
									v90 = v55
								case 1, 2:
									v90 = v61
								default:
									v90 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v90
							}
						}
					}
				case 3:
					v64 = F_TS_phrase_execute(m, l0, l1, l2, l3, int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						switch v64 {
						case 0, 1:
							v90 = v64
						case 2:
							v90 = l2 & int32(2)
						default:
							v90 = int32(0)
						}
						m.G0 = v10 + int32(16)
						return v90
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
						F_errmsg_internal(m, int32(_a_F_TS_execute_recurse_0), v10)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_TS_execute_recurse_1), int32(1969), int32(_a_F_TS_execute_recurse_2))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
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
	}
}
func F_TS_phrase_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	F_check_stack_depth(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_TS_phrase_execute[0]))
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v16 + int32(48)
	return v401
L8:
	;
	v29 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, l4)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	switch v31 - int32(1) {
	case 0:
		goto L18
	case 1, 3:
		goto L17
	case 2:
		goto L16
	default:
		goto L13
	}
L11:
	;
	v401 = v29
	goto L7
L12:
	;
	v401 = int32(0)
	goto L7
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L113
	}
L14:
	;
	if v312 != 0 {
		goto L104
	} else {
		goto L105
	}
L15:
	;
	v340 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v340)
	v401 = v340
	goto L7
L16:
	;
	v261 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v261
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v275 = F_TS_phrase_execute(m, l0+v269*int32(12), l1, l2, l3, v16+int32(32))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L85
	}
L17:
	;
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v60
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = F_TS_phrase_execute(m, l0+v68*int32(12), l1, l2, l3, v16+int32(32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L28
	}
L18:
	;
	v34 = int32(1)
	if l2&v34 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v37)
	v401 = v34
	goto L7
L20:
	;
	goto L21
L21:
	;
	v41 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, l4)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(0) < v45 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	switch v41 {
	case 0:
		goto L15
	case 1:
		goto L22
	case 2:
		v401 = v41
		goto L7
	default:
		goto L12
	}
L24:
	;
	v51 = (v44 ^ int32(-1)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v51)
	v401 = int32(1)
	goto L7
L25:
	;
	goto L26
L26:
	;
	if v44&int32(1) == int32(0) {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v57)
	v401 = v57
	goto L7
L28:
	;
	if v74 == int32(0) {
		v401 = v6
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v82 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, v16+int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v82 == int32(0) {
		v401 = v6
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v86 = int32(2)
	if base.B2i32(v74 == v86)|base.B2i32(v82 == v86) != 0 {
		v401 = v86
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v92 == int32(4) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)))
	v122 = int32(1)
	v124 = int32(0)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+37)))
	if base.B2i32(v121&v122 == v124)|base.B2i32(v126 != v122) == v124 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v110
	v119 = v113
	v120 = v114
	goto L33
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v97 = v95 + v96
	if l4 == int32(0) {
		v119 = v97
		v120 = v6
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v103 < v102 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v110 = v97 + v100
	v113 = v97
	v114 = v6
	goto L34
L39:
	;
	v105 = v102
	goto L41
L40:
	;
	v105 = v103
	goto L41
L41:
	;
	v106 = v105 - v103
	v107 = v105 - v102
	if l4 == int32(0) {
		v119 = v107
		v120 = v106
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v110 = v105
	v113 = v107
	v114 = v106
	goto L34
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v137 = int32(0)
	v141 = v132
	v142 = v137
	v145 = v137
	goto L46
L44:
	;
	goto L45
L45:
	;
	if v126 != 0 {
		goto L73
	} else {
		goto L74
	}
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v158 = v142
	v161 = v145
	goto L48
L48:
	;
	if v141 <= v161 {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	if l4 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L50:
	;
	if v206 <= int32(0) {
		v158 = v208
		v161 = v209
		goto L48
	} else {
		goto L65
	}
L51:
	;
	v206 = v188
	v208 = v158 + int32(1)
	v209 = v161
	goto L50
L52:
	;
	v200 = int32(1)
	v206 = v198
	v208 = v158 + v200
	v209 = v161 + v200
	goto L50
L53:
	;
	if v181 != v188 {
		goto L51
	} else {
		goto L64
	}
L54:
	;
	v192 = int32(2147483647)
	if v178 == v192 {
		v198 = v192
		goto L52
	} else {
		goto L63
	}
L55:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152+v158<<(uint(int32(1))%32)))))
	v188 = v120 + v185&int32(_a_F_TS_phrase_execute_0)
	if v188 <= v181 {
		goto L53
	} else {
		goto L62
	}
L56:
	;
	if v158 < v153 {
		v181 = int32(2147483647)
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154+v161<<(uint(int32(1))%32)))))
	v178 = v119 + v175&int32(_a_F_TS_phrase_execute_0)
	if v153 <= v158 {
		goto L54
	} else {
		goto L61
	}
L59:
	;
	if l4 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	v401 = int32(1)
	goto L7
L61:
	;
	v181 = v178
	goto L55
L62:
	;
	v206 = v181
	v208 = v158
	v209 = v161 + int32(1)
	goto L50
L63:
	;
	v206 = v178
	v208 = v158
	v209 = v161 + int32(1)
	goto L50
L64:
	;
	v198 = v181
	goto L52
L65:
	;
	goto L49
L66:
	;
	v401 = int32(1)
	goto L7
L67:
	;
	goto L68
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v215 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v218 = F_palloc(m, (v132+v133)<<(uint(int32(1))%32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	v223 = v215
	goto L71
L71:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v225 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v224 + v225
	*(*uint16)(unsafe.Add(mBase, uint32(v223+v224<<(uint(v225)%32)))) = uint16(v206)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v141 = v232
	v142 = v208
	v145 = v209
	goto L46
L72:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(v220)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v218
	v223 = v218
	goto L71
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v239 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), int32(2), v119, v120, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v121&int32(1) != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v401 = v239
	goto L7
L77:
	;
	v249 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), int32(1), v119, v120, v241)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v241 < v256 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v401 = v249
	goto L7
L81:
	;
	v258 = v241
	goto L83
L82:
	;
	v258 = v256
	goto L83
L83:
	;
	v259 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), int32(4), v119, v120, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v401 = v259
	goto L7
L85:
	;
	v281 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, v16+int32(16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v275|v281 == int32(0) {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	v286 = int32(2)
	if base.B2i32(v275 == v286)|base.B2i32(v281 == v286) != 0 {
		v401 = v286
		goto L7
	} else {
		goto L88
	}
L88:
	;
	if v275 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = int32(0)
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v281 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	if v300 < v301 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v300 = v296
	goto L92
L94:
	;
	goto L95
L95:
	;
	v297 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v297
	v300 = v297
	goto L92
L96:
	;
	v303 = v301
	goto L98
L97:
	;
	v303 = v300
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v303
	v305 = v303 - v300
	v306 = v303 - v301
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)))
	v308 = int32(1)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+37)))
	if base.B2i32(v307&v308 == int32(0))|base.B2i32(v312 != v308) != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v321 < v322 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v324 = v321
	goto L102
L101:
	;
	v324 = v322
	goto L102
L102:
	;
	v325 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), int32(4), v306, v305, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L15
L104:
	;
	v343 = int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v350 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), v343, v306, v305, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v354 = int32(1)
	if v307&v354 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v352 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v352)
	v401 = v343
	goto L7
L108:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v363 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), int32(2), v306, v305, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v375 = F_TS_phrase_output(m, l4, v16+int32(32), v16+int32(16), int32(7), v306, v305, v372+v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v365 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v365)
	v401 = v354
	goto L7
L112:
	;
	v401 = v375
	goto L7
L113:
	;
	v381 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v381
	F_errmsg_internal(m, int32(_a_F_TS_phrase_execute_1), v16)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_TS_phrase_execute_2), int32(1837), int32(_a_F_TS_phrase_execute_3))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_error_commit_ts_disabled(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_errcode(m, int32(325))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_error_commit_ts_disabled_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_error_commit_ts_disabled[0])))
				if v20 == int32(1) {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_error_commit_ts_disabled[1]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+316))
					v28 = base.B2i32(v26 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _c_F_error_commit_ts_disabled[0])) = uint8(v28)
					v30 = v28
				} else {
					v30 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_error_commit_ts_disabled_1)
				if v30 != 0 {
					v35 = int32(_a_F_error_commit_ts_disabled_2)
				} else {
					v35 = int32(_a_F_error_commit_ts_disabled_3)
				}
				F_errhint(m, v35, v5)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_error_commit_ts_disabled_4), int32(390), int32(_a_F_error_commit_ts_disabled_5))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
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
func F_get_ts_parser_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_defGetQualifiedName(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(2281)
		v18 = int32(1)
		switch l1 - int32(5) {
		case 0:
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(9796820404457)
			v30 = int32(3)
			v32 = int32(2281)
			v33 = v30
		case 1:
			v32 = int32(2278)
			v33 = v18
		case 2:
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(15526306777321)
			v30 = int32(3)
			v32 = int32(2281)
			v33 = v30
		case 3:
			v30 = v18
			v32 = int32(2281)
			v33 = v30
		default:
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(23)
			v30 = int32(2)
			v32 = int32(2281)
			v33 = v30
		}
		v35 = v9 + int32(20)
		v37 = F_LookupFuncName(m, v11, v33, v35, int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = F_get_func_rettype(m, v37)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				if v39 != v32 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v50 = F_func_signature_string(m, v11, v33, int32(0), v35)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = F_format_type_be(m, v32)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v50
									F_errmsg(m, int32(_a_F_get_ts_parser_func_0), v9)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_ts_parser_func_1), int32(126), int32(_a_F_get_ts_parser_func_2))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
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
				} else {
					m.G0 = v9 + int32(32)
					return v37
				}
			}
		}
	}
}
func F_ts_headline_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_DirectFunctionCall3Coll(m, int32(1174), int32(0), v4, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_ts_headline_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_DirectFunctionCall3Coll(m, int32(1173), int32(0), v4, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_ts_match_tq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_DirectFunctionCall1Coll(m, int32(1525), v2, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_DirectFunctionCall2Coll(m, int32(1523), v2, v15, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v15)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v21 != v5 {
						F_pfree(m, v5)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v17 != int32(0))
						}
					} else {
						return base.B2i32(v17 != int32(0))
					}
				}
			}
		}
	}
}
func F_ts_rank_tt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 float32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_calc_rank(m, int32(_a_F_ts_rank_tt_0), v7, v11, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v15 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v13)
						}
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v19 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				} else {
					return base.I32_reinterpret_f32(v13)
				}
			}
		}
	}
}
func F_ts_rank_ttf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = F_calc_rank(m, int32(_a_F_ts_rank_ttf_0), v7, v11, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v15 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v13)
						}
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v19 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				} else {
					return base.I32_reinterpret_f32(v13)
				}
			}
		}
	}
}
func F_ts_token_type_byname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = F_textToQualifiedNameList(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v18 = F_get_ts_parser_oid(m, v15, int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_tt_setup_firstcall(m, v13, l0, v18)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
							v25 = F_tt_process_call(m, v24)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								if v25 != 0 {
									v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
									*(*int64)(unsafe.Add(mBase, uint32(v24))) = v27 + int64(1)
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
									return v25
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
										v40 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
										return v25
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		v25 = F_tt_process_call(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v25 != 0 {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
				*(*int64)(unsafe.Add(mBase, uint32(v24))) = v27 + int64(1)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
				return v25
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
					return v25
				}
			}
		}
	}
}
