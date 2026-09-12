package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CteScanNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+132))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	F_tuplestore_select_read_pointer(m, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v17*int32(24))+4)))
		if v7 == int32(1) {
			if v21 != 0 {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
				if v42 != 0 {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
					m.T0[v71].(func(*base.Module, int32))(m, v15)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v77 = v15
						return v77
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
					if v44 != 0 {
						F_ExecReScan(m, v43)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
							v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 != 0 {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
									if v50&int32(2) == int32(0) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										F_tuplestore_select_read_pointer(m, v9, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_tuplestore_puttupleslot(m, v9, v48)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
												m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v77 = v15
													return v77
												}
											}
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
										return int32(0)
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
									return int32(0)
								}
							}
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
						v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 != 0 {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
								if v50&int32(2) == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									F_tuplestore_select_read_pointer(m, v9, v60)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_tuplestore_puttupleslot(m, v9, v48)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
											m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v77 = v15
												return v77
											}
										}
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
									return int32(0)
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
								return int32(0)
							}
						}
					}
				}
			} else {
				v33 = int32(1)
				v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 != 0 {
						v77 = v15
						return v77
					} else {
						if v7 != int32(1) {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
							m.T0[v71].(func(*base.Module, int32))(m, v15)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v77 = v15
								return v77
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
							if v42 != 0 {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = v15
									return v77
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
								if v44 != 0 {
									F_ExecReScan(m, v43)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
												if v50&int32(2) == int32(0) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v60)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v48)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
															m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return int32(0)
															} else {
																v77 = v15
																return v77
															}
														}
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										}
									}
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
									v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 != 0 {
											v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
											if v50&int32(2) == int32(0) {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												F_tuplestore_select_read_pointer(m, v9, v60)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													F_tuplestore_puttupleslot(m, v9, v48)
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return int32(0)
													} else {
														v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
														m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															v77 = v15
															return v77
														}
													}
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v56 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v21 == int32(0) {
				if v21 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
					if v42 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
						m.T0[v71].(func(*base.Module, int32))(m, v15)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v77 = v15
							return v77
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
						if v44 != 0 {
							F_ExecReScan(m, v43)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
								v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 != 0 {
										v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
										if v50&int32(2) == int32(0) {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
											F_tuplestore_select_read_pointer(m, v9, v60)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												F_tuplestore_puttupleslot(m, v9, v48)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
													m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														v77 = v15
														return v77
													}
												}
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v56 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
											return int32(0)
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
										return int32(0)
									}
								}
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
							v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 != 0 {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
									if v50&int32(2) == int32(0) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										F_tuplestore_select_read_pointer(m, v9, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_tuplestore_puttupleslot(m, v9, v48)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
												m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v77 = v15
													return v77
												}
											}
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
										return int32(0)
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
									return int32(0)
								}
							}
						}
					}
				} else {
					v33 = int32(1)
					v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v77 = v15
							return v77
						} else {
							if v7 != int32(1) {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = v15
									return v77
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
								if v42 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
									m.T0[v71].(func(*base.Module, int32))(m, v15)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v77 = v15
										return v77
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
									if v44 != 0 {
										F_ExecReScan(m, v43)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
											v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 != 0 {
													v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
													if v50&int32(2) == int32(0) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
														F_tuplestore_select_read_pointer(m, v9, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_tuplestore_puttupleslot(m, v9, v48)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int32(0)
																} else {
																	v77 = v15
																	return v77
																}
															}
														}
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
														return int32(0)
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
												if v50&int32(2) == int32(0) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v60)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v48)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
															m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return int32(0)
															} else {
																v77 = v15
																return v77
															}
														}
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+136)))
				if v27 != 0 {
					v33 = int32(1)
					v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 != 0 {
							v77 = v15
							return v77
						} else {
							if v7 != int32(1) {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = v15
									return v77
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
								if v42 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
									m.T0[v71].(func(*base.Module, int32))(m, v15)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v77 = v15
										return v77
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
									if v44 != 0 {
										F_ExecReScan(m, v43)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
											v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 != 0 {
													v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
													if v50&int32(2) == int32(0) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
														F_tuplestore_select_read_pointer(m, v9, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_tuplestore_puttupleslot(m, v9, v48)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int32(0)
																} else {
																	v77 = v15
																	return v77
																}
															}
														}
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
														return int32(0)
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
										v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 != 0 {
												v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
												if v50&int32(2) == int32(0) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v60)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v48)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
															m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return int32(0)
															} else {
																v77 = v15
																return v77
															}
														}
													}
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
													return int32(0)
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v56 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v28 = int32(0)
					v30 = F_tuplestore_advance(m, v9, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v33 = int32(1)
							v36 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v33), v33, v15)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 != 0 {
									v77 = v15
									return v77
								} else {
									if v7 != int32(1) {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
										m.T0[v71].(func(*base.Module, int32))(m, v15)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v77 = v15
											return v77
										}
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+136)))
										if v42 != 0 {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
											m.T0[v71].(func(*base.Module, int32))(m, v15)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v77 = v15
												return v77
											}
										} else {
											v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
											if v44 != 0 {
												F_ExecReScan(m, v43)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
													v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														if v48 != 0 {
															v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
															if v50&int32(2) == int32(0) {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
																F_tuplestore_select_read_pointer(m, v9, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return int32(0)
																} else {
																	F_tuplestore_puttupleslot(m, v9, v48)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return int32(0)
																	} else {
																		v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																		m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																		mBase = m.M
																		v68 = m.ExcPending
																		if v68 != 0 {
																			return int32(0)
																		} else {
																			v77 = v15
																			return v77
																		}
																	}
																}
															} else {
																v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
																v56 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
																return int32(0)
															}
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															v56 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
															return int32(0)
														}
													}
												}
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
												v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v43)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													if v48 != 0 {
														v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
														if v50&int32(2) == int32(0) {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
															F_tuplestore_select_read_pointer(m, v9, v60)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return int32(0)
															} else {
																F_tuplestore_puttupleslot(m, v9, v48)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return int32(0)
																} else {
																	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
																	m.T0[v66].(func(*base.Module, int32, int32))(m, v15, v48)
																	mBase = m.M
																	v68 = m.ExcPending
																	if v68 != 0 {
																		return int32(0)
																	} else {
																		v77 = v15
																		return v77
																	}
																}
															}
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															v56 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
															return int32(0)
														}
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v55)+136)) = uint8(v56)
														return int32(0)
													}
												}
											}
										}
									}
								}
							}
						} else {
							v77 = v28
							return v77
						}
					}
				}
			}
		}
	}
}
